# 100 Lamps Problem

## Given
- There is 100 lamps labeled with number
- The lamps are turned off in the beginning
- First trip, all lamps turned on
- Second trip, all lamps with number multiplying of two switched to off
- Next trip, all lamps with number multiplying of three switched to off
- And so on, until the 100th step

## Problem
How many lamps turned on in the end?

## Solution
A lamp only switched in i<sup>th</sup> step if divided by lamp number. For example:
- 15 switched in 1st, 3rd, 5th, 15th (4 times)
- 60 switched in 1st, 2nd, 3rd, 4th, 5th, 6th, 10th, 12th, 15th, 20th, 30th, 60th (12 times)

It shows that the lamp will turned off in the end if the lamp switched in an even number of steps. How we get the odd number, so the lamp turned off?

Lamp with squared number switched in an odd number of steps. So, in the end it will turned on.

So the provided solution is find the squared number less than number of lamps. So it will be the number of iteration which represents how much the lamps remain on, then square it. So it will produce the lamps number remains on.
