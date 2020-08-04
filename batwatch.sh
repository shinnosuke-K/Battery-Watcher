#!/bin/bash

array=()

caps=$(ioreg -l |
grep -v 'Apple' |
grep -v 'BatteryData' |
grep -e 'MaxCapacity' -e 'DesignCapacity' -e 'CurrentCapacity' |
awk '{print $3 $5}')

echo $caps
# rate=$((${array[2]}/${array[0]}))
# array=($rate ${array[@]})
# echo $array
