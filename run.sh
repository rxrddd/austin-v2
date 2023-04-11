docker run --rm -d -p 8000:8000 --network deply_software_network  -v /www/wwwroot/austin-v2/app/mgr/configs:/data/conf mgr

docker run --rm -d -p 7000:7000 -p 8001:8000 --network deply_software_network  -v /www/wwwroot/austin-v2/app/msgpusher/configs:/data/conf msgpusher

docker run --rm -d -p 10000:8000 --network deply_software_network  -v /www/wwwroot/austin-v2/app/msgpusher-manager/configs:/data/conf msgpusher-manager

docker run --rm -d --network deply_software_network  -v /www/wwwroot/austin-v2/app/msgpusher-worker/configs:/data/conf msgpusher-worker