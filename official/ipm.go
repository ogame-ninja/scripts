origin = "1:2:3"                                         // Planet to use
target = "4:5:6"                                         // Target coordinate
nbr = 20                                                 // Number of missiles to build and send
constructionTime = 17                                    // Time to build 1 missile (secs)

for i = 1; i <= 30; i++ {                                // Repeat the attack 30 times
    BuildDefense(origin, INTERPLANETARYMISSILES, nbr)    // Build missiles
    SleepSec((constructionTime+1) * nbr)                 // Wait for missiles to be built
    duration, err = SendIPM(planet.ID, target, nbr, 0)   // Send missiles
    Print(duration, err)                                 // Print error if any
}
