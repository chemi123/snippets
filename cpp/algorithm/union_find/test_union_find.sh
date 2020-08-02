
for i in {1..100}
do
python3 union_find_dataset.py > dataset
./a.out < dataset > /dev/null
#if [ $? -eq 139 ]; then
#  echo $?
#fi
done