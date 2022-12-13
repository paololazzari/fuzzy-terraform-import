sudo docker rm -f testcontainer >/dev/null
sudo docker run --name testcontainer -d -p 5000:5000 -v /mnt/c/Users/Paolo/Documents/GitHub/fuzzy-terraform-import/test:/src/ testimage
sudo docker exec testcontainer /test/test_wrapper.sh
sudo docker exec testcontainer go test -v ./...