
echo "Download project dependences"
go get ./...

echo "Uncompress files"
tar xvzf *.tgz >> /dev/null

for FILE in $(ls | grep .csv);
do
   echo "Removing header from $FILE";
   vim -u NONE +'1d' +wq! $FILE;
done

echo "Creating postgres container"
docker-compose up -d

echo "Start importer"
sleep 5s
go run importer.go


echo "Remove imported files"
rm *.csv
