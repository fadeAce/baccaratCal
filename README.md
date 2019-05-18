# baccaratCal

calculate baccarat win rate and lose rate by computation theory without distribution preset or any knowledge for it.

#### how to run it ?
    // by go run env
    go build calculation
    
    // run it by redirect standrad output
    ./calculation 1>result.log
    
#### the result should be like 
```$xslt
...
11   11   13   13   3.333738451068494e-05  atimes  0 btimes  1
11   12   13   13   6.882556802205923e-05  atimes  0 btimes  1
11   13   13   13   6.452397002068053e-05  atimes  1 btimes  2
12   12   13   13   3.333738451068494e-05  atimes  0 btimes  1
12   13   13   13   6.452397002068053e-05  atimes  1 btimes  2
13   13   13   13   2.9237423915620866e-05  atimes  2 btimes  3
all types sum to  0.9999999999999297
the root types of base card range is  8281
the result win is  0.4585974226327612  chargerlose  0.4462466093436007  chargerdie  0.09515596802364121
sum is  1.000000000000003
```

#### explanation
For reason why sum is beyond 1 , you can personally remake data structure with more precise float, but it's fair enough for precision at 10^(-14) 
as mentioned at [WIKIPEDIA](https://zh.wikipedia.org/wiki/%E7%99%BE%E5%AE%B6%E6%A8%82)