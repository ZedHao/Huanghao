for line in $(cat keyword.conf)
do
   grep  -r $1  -nFe ${line} 
done
