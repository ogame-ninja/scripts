/// -------------------------------------
/// Defense Builder v.1.0 by @toverux
///
/// This script is able to build your defense on each planet/moon according to a
/// single global ratio for each defense installation, and a multiplying factor
/// for each planet/moon.
/// It does NOT handle ABM/IPMs yet.
///
/// IMPORTANT NOTICE:
/// Contrarily to other scripts out there, this program uses the bot's brain
/// construction queue, meaning the brain will then take care of the importing/
/// running build orders. This is better for those of us who have specialized
/// planets with not enough of the three resources to build defense at any
/// moment.
/// Because the bot's queue can't be queried by scripts (yet), this script can't
/// run continuously, because it doesn't know whether a build order for a given
/// defense building has already been programmed or not.
/// For this reason, the script must be executed manually once in a while, by
/// you. It will ERASE the bot's construction queue for ALL planets/moons and
/// reprogram it.
/// The script takes care about querying your shipyard's queue before ordering
/// more buildings, though, so if you need 30 LLs total, have already built 10,
/// and the shipyard's building 10, the script will only order 10 more.
///
/// The script will warn you if a building can't be built due to unmet
/// requirements (eg. tech or shipyard level too low).
///
/// Future work:
///  - Auto build shipyard when needed
///  - Build IPMs and ABMs as well
///  - When the scripting API allows it, stop erasing the queues before adding
///    buildings to it
///  - (depends on step above) Long-running version for automatic (and not
///    one-shot) building to maintain defenses, notably after attacks.
///
/// Changelog:
///  - v.1.0: first useable version
/// -------------------------------------

/// ---- [ CONFIGURATION ] ---- ///

//=> Whether to run a simulation or not. Use `true` to simulate and the script
//   will just tell you what's missing on each celestial, without touching your
//   building queue. Use `false` to enable the construction ordering logic.
dryRun = true

//=> The default multiplication factor for all planets and moons.
//   Use `defaultFactor = nil` to skip all planets that do not have an explicit
//   multiplying factor set.
defaultFactor = 1

//=> A list of explicit multiplication factors for planet/moons.
//   Setting a value for a planet/moon will override `defaultFactor` above.
//   Your planet/moons MUST have an unique name in order to be used in this
//   list.
factors = {
    "YOUR PLANET NAME": 4, // this planet will have 4x the amount specified in `wanted`
    "YOUR MOON NAME": 0.6, // this moon will have 0.6x the amount specified in `wanted`
}

//=> The "ratios" for each defense building. Each amount here will be multiplied
//   by `defaultFactor` or the per-planet/moon override specified in `factors`.
//   You can set the amount to 0 or remove a line entirely if you don't want a
//   kind of building at all.
wanted = {
    ROCKETLAUNCHER: 1000,
    LIGHTLASER: 500,
    HEAVYLASER: 100,
    GAUSSCANNON: 20,
    IONCANNON: 50,
    PLASMATURRET: 10,
    SMALLSHIELDDOME: 1,
    LARGESHIELDDOME: 1
}

/// --- [ START OF THE SCRIPT (do not change code below - unless you know what you're doing) ] ---- ///

func MaybeAbort(err) {
    if err != nil && err != 0 {
        LogError(err)
        Terminate()
    }
}

func AddToQueue(celestial, id, quantity) {
    if quantity == 0 {
        return
    }

    supplies, facilities, _, _, researches, err = GetTechs(celestial.ID)
    MaybeAbort(err)

    canBuild = IsAvailable(id, celestial.ID, supplies, facilities, researches, 0)

    if !canBuild {
        requirements = GetRequirements(id)

        LogError(celestial.Name + ": cannot build " + id + " x" + quantity + " yet, please ensure requirements are met: " + requirements)

        return
    }

    LogWarn(celestial.Name + ": " + (dryRun ? "missing: " : "adding to build queue: ") + id + " x" + quantity)

    if dryRun {
        return
    }

    err = AddItemToQueue(celestial.ID, id, quantity)
    MaybeAbort(err)
}

func GetProductionForId(production, id) {
    amount = 0

    for item in production {
        if (item.ID == id) {
            amount += item.Nbr
        }
    }

    return amount
}

func GetMissing(id, factor, built, production) {
    wantedTotal = Round(wanted[id] ?? 0 * factor)
    missing = wantedTotal - built - GetProductionForId(production, id)

    return Max(0, missing)
}

celestials, _ = GetCelestials()

areQueuesErased = false

for celestial in celestials {
    factor = factors[celestial.Name] ?? defaultFactor

    if (factor == nil) {
        LogWarn(celestial.Name + ": no factor defined, ignoring.")
        continue
    }

    LogInfo(celestial.Name + ": checking defenses (x" + factor + ")...")

    defense, _ = celestial.GetDefense()
    production, _ = celestial.GetProduction()

    missingRocketLaunchers = GetMissing(ROCKETLAUNCHER, factor, defense.RocketLauncher, production)
    missingLightLasers = GetMissing(LIGHTLASER, factor, defense.LightLaser, production)
    missingHeavyLasers = GetMissing(HEAVYLASER, factor, defense.HeavyLaser, production)
    missingGaussCannons = GetMissing(GAUSSCANNON, factor, defense.GaussCannon, production)
    missingIonCannons = GetMissing(IONCANNON, factor, defense.IonCannon, production)
    missingPlasmaTurrets = GetMissing(PLASMATURRET, factor, defense.PlasmaTurret, production)

    missingSmallShieldDomes = GetMissing(SMALLSHIELDDOME, 1, defense.SmallShieldDome, production)
    missingLargeShieldDomes = GetMissing(LARGESHIELDDOME, 1, defense.LargeShieldDome, production)

    missingUnits = missingRocketLaunchers + missingLightLasers + missingHeavyLasers + missingGaussCannons + missingIonCannons + missingPlasmaTurrets + missingSmallShieldDomes + missingLargeShieldDomes

    if missingUnits > 0 && !areQueuesErased && !dryRun {
        LogWarn("Erasing ALL construction queues...")
        ClearAllConstructionQueues()
        areQueuesErased = true
    }

    if missingUnits == 0 {
        LogInfo(celestial.Name + ": already has all its defenses built or programmed in the shipyard.")
    }

    AddToQueue(celestial, ROCKETLAUNCHER, missingRocketLaunchers)
    AddToQueue(celestial, LIGHTLASER, missingLightLasers)
    AddToQueue(celestial, HEAVYLASER, missingHeavyLasers)
    AddToQueue(celestial, GAUSSCANNON, missingGaussCannons)
    AddToQueue(celestial, IONCANNON, missingIonCannons)
    AddToQueue(celestial, PLASMATURRET, missingPlasmaTurrets)

    AddToQueue(celestial, SMALLSHIELDDOME, missingSmallShieldDomes)
    AddToQueue(celestial, LARGESHIELDDOME, missingLargeShieldDomes)

    SleepSec(2)
}

LogInfo("Done.")