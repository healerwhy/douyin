sql2pb -go_package ./pb -package pb -host 120.79.222.123 -user healer -password healer000. -port 3306 \
        -schema douyin -service_name VideoService -table video,  > VideoService.proto