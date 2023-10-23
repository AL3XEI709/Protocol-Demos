echo open in clash
git config --global http.proxy socks5 127.0.0.1:7890
git config --global https.proxy socks5 127.0.0.1:7890
git config --global http.proxy 127.0.0.1:7890
git config --global https.proxy 127.0.0.1:7890
git add -A
git commit -am "newest" 
git push