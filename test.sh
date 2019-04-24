wget https://dl.google.com/go/go1.12.3.linux-amd64.tar.gz 
tar vzfx go1.12.3.linux-amd64.tar.gz 
mv go /usr/bin/
#export PATH=/usr/local/go/bin:/usr/bin
echo "export PATH=$PATH:/usr/bin/go/bin" >> /root/.bash_profile
echo "export GOPATH=/go" >> /root/.bash_profile
spurce /root/.bash_profile
rm -f go1.12.3.linux-amd64.tar.gz
touch aaa.txt
