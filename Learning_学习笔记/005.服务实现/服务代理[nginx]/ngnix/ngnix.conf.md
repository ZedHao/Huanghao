作业帮的ngnix.conf
```

 user  homework;
 worker_processes  2;
 worker_cpu_affinity 00000001 00000010;
 worker_rlimit_nofile 8192;
 
 error_log   "/home/homework/log/webserver/error_log"   notice;
 pid         "/home/homework/log/nginx.pid";
 
 events {
     use epoll;
     worker_connections  8192; 
 }
 
 http {
     include       mime.types; 
     #include       upstream.conf;
     default_type  application/octet-stream;
 
     log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                       '$status $body_bytes_sent "$http_referer" "$http_cookie" "$http_user_agent" '
                       '$request_time $logid $remote_addr $server_addr $upstream_addr $host '
                       '"$http_x_forwarded_for" $product $subsys $tracecode $msec';
 
     access_log  "/home/homework/log/webserver/access_log"  main;
 
     client_body_temp_path /home/homework/webserver/cache/client_body;
     fastcgi_temp_path /home/homework/webserver/cache/fastcgi;
     proxy_temp_path /home/homework/webserver/cache/proxy;
     uwsgi_temp_path /home/homework/webserver/cache/uwsgi;
     scgi_temp_path /home/homework/webserver/cache/scgi;
 
     server_names_hash_bucket_size 128;
     client_header_buffer_size 16k;
     large_client_header_buffers 4 128k;
     client_max_body_size 50m;
     client_body_buffer_size 1024k;
 
     sendfile        on;
     tcp_nopush      on;
     tcp_nodelay     on;
 
     fastcgi_connect_timeout 5;
     fastcgi_send_timeout 300;
     fastcgi_read_timeout 300;
     fastcgi_buffer_size 128k;
     fastcgi_buffers 4 256k;
     fastcgi_busy_buffers_size 512k;
     fastcgi_temp_file_write_size 512k;
     fastcgi_intercept_errors on;
 
     keepalive_timeout  0;
     #keepalive_timeout  65;
 
     gzip on;
     gzip_min_length 1k;
     gzip_buffers 4 16k;
     gzip_http_version 1.0;
     gzip_comp_level 2;
     gzip_types text/plain application/x-javascript text/css application/xml;
     gzip_vary on;
 
     policy_frame off; 
     policy_path "/home/homework/webserver/conf/";
     policy_dtdname policy.dtd;
     policy_docname policy.xml;
     policy_logpath "/home/homework/log/webserver";
     policy_logfile policy.;
     policy_logsize 1600;
     policy_loglevel 4; 
 
     uninitialized_variable_warn off;
     proxy_connect_timeout 15s;
     proxy_read_timeout 10s;
     proxy_send_timeout 10s;
     proxy_buffer_size 64k;
     proxy_buffers 4 64k;
     proxy_busy_buffers_size 128k;
     proxy_temp_file_write_size 128k;
     proxy_set_header X_BD_LOGID $logid;
     proxy_set_header LOGID $logid;
     proxy_set_header Host $http_host;
     proxy_set_header X-Real-IP $remote_addr;
     proxy_set_header REMOTE-HOST $remote_addr;
     proxy_set_header X-Forwarded-For $remote_addr;
     proxy_set_header CLIENTIP $remote_addr;
     proxy_http_version 1.1;
 
     port_in_redirect off;
     set_real_ip_from 192.168.0.0/24;
     real_ip_header CLIENTIP;
     underscores_in_headers on;
     include vhost/php.conf;
 }    
```

开发机的ngnix.conf

```



#nginx进程数，建议设置为等于CPU总核心数
worker_rlimit_nofile 204800;
#全局错误日志定义类型，[ debug | info | notice | warn | error | crit ]
error_log /var/log/nginx/error.log;
#工作模式及连接数上限
events {
        #参考事件模型，use [ kqueue | rtsig | epoll | /dev/poll | select | poll ];
        #epoll模型是Linux 2.6以上版本内核中的高性能网络I/O模型，如果跑在FreeBSD上面，就用kqueue模型
        worker_connections  102400;
}
#设定http服务器，利用它的反向代理功能提供负载均衡支持
http {
        #文件扩展名与文件类型映射表
        #首先，我们要了解浏览器是如何处理内容的。在浏览器中显示的内容有 HTML、有 XML、有 GIF、还有 #Flash ……那么，浏览器是如何区分它们，决定什么内容用什么形式来显示呢？答案是   #MIMEType，也就是该资源的媒体类型。
        include       mime.types;
        #默认文件类型
        default_type  application/octet-stream;
        #设定日志格式
        #log_format  main  '$idXXXX\t$remote_addr\t$remote_user\t$time_local\t$http_host\t$request\t'
        #             '$status\t$body_bytes_sent\t$http_referer\t'
        #             '$http_user_agent\t$http_x_forwarded_for\t$request_time\t$upstream_response_time\t$userid';
                        "$request_time\t$status\t$body_bytes_sent\t'$http_referer'\t"

    #不可见字符分隔日志格式
    #include other_log_format.conf;
    #实时日志收集json格式日志
    #include json_log_format.conf;

        #日志流格式
                      "$request_time\t$status\t$body_bytes_sent\t'$http_referer'\t"
                      "'$http_user_agent'\t'$http_x_forwarded_for'\t$upstream_addr\t$upstream_response_time\t3";

        #成功日志
        access_log  /var/log/nginx/access.log  main;
 #access_log syslog:local6:notice:log1.op.XXXXdns.org:514:nginx-main-log main;

        #指定 nginx 是否调用 sendfile 函数（zero copy 方式）来输出文件，对于普通应用，
        sendfile        on;
        #长连接超时时间，单位是秒
        keepalive_timeout  60;
        #服务器名称哈希表的最大值(默认512)[hash%size]
        server_names_hash_max_size 1024;
        #服务器名字的hash表大小
        server_names_hash_bucket_size 256;
        #客户请求头缓冲大小
        client_header_buffer_size 4k;
        #如果header过大，它会使用large_client_header_buffers来读取
        large_client_header_buffers 4 256k;
        client_header_timeout  1m;
        client_body_timeout    1m;
        send_timeout           1m;
        #防止网络阻塞
        tcp_nopush     on;
        tcp_nodelay    on;
        #允许客户端请求的最大单文件字节数
        client_max_body_size 50m;

        #缓冲区代理缓冲用户端请求的最大字节数
        client_body_buffer_size 50m;

        #nginx跟后端服务器连接超时时间(代理连接超时)
        proxy_connect_timeout 5;

        #后端服务器数据回传时间(代理发送超时)
        proxy_send_timeout 15;

        #连接成功后，后端服务器响应时间(代理接收超时)
        proxy_read_timeout 15;

        #设置代理服务器（nginx）保存用户头信息的缓冲区大小
        proxy_buffer_size 4k;

        #proxy_buffers缓冲区，网页平均在32k以下的话，这样设置
        proxy_buffers 8 32k;

        #高负荷下缓冲大小（proxy_buffers*2）
        proxy_busy_buffers_size 64k;
        proxy_temp_file_write_size 64k;
        proxy_intercept_errors  on;
        #客户端放弃请求，nginx也放弃对后端的请求
        #proxy_ignore_client_abort on;

    #代理缓存头信息最大长度[设置头部哈希表的最大值，不能小于你后端服务器设置的头部总数]
    proxy_headers_hash_max_size 512;
    #设置头部哈希表大小(默认64)[这将限制头部字段名称的长度大小，如果你使用超过64个字符的头部名可以加大这个值。]
    proxy_headers_hash_bucket_size 256;

    #变量哈希表的最大值(默认值)
    variables_hash_max_size 512;
    #为变量哈希表制定关键字栏的大小(默认64)
    variables_hash_bucket_size 128;

        #开启gzip压缩输出
        gzip on;
        #最小压缩文件大小
        gzip_min_length 1k;
        #压缩缓冲区
        gzip_buffers     4 16k;
        #压缩等级
        gzip_comp_level 9;
        #压缩版本（默认1.1，前端如果是squid2.5请使用1.0）
        gzip_http_version 1.0;
        #压缩类型，默认就已经包含textml
        gzip_vary on;

        #map模块使用
        map_hash_max_size 102400;
        map_hash_bucket_size  256;

        #Tengine Config
        #concat on;
        #trim on;
        #trim_css off;
        #trim_js off;
        server_tokens off;
        #footer "<!-- $remote_addr $server_addr $upstream_addr -->";

        #rewrite_log on;
        fastcgi_intercept_errors on;
    #include other config file
        include ../conf.d/*.conf;
    #包含一些特殊站点的配置文件,此目录下文件暂时不包含在git自动管理过程中
    include ../special/*.conf;

    #屏蔽不加主机域名的默认请求
    #server {
        #   listen *:80 default;
        #   server_name _ "";
        #   return 444;
    #}

    #ceshi by dongange 2016-01-07

    #Nginx状态监测模块配置
#    req_status_zone server "$server_name,$server_addr:$server_port" 10M;
 #   req_status server;
    server {
          listen       80 default_server;
        server_name  39.105.184.207;
        #root         /usr/share/nginx/html;
        #root /root/hao/haoChat/Applications/Chat/Web/;
        #root /root/hao/haoMall/shopMall/protected/view/backend/;
        #root /root/hao/haoMall/shopMall/;
        root /root/hao;
             location / {
        root /root/hao;
           index index.php;
          fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        }
         location ~ \.php$ {
                root /root/hao;
                fastcgi_connect_timeout  180;
                fastcgi_read_timeout    600;
                fastcgi_send_timeout    600;
                fastcgi_pass 127.0.0.1:9000;
                #fastcgi_pass   unix:/tmp/php-cgi.sock;
                fastcgi_index index.php;
                fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
                include fastcgi_params;
        }

        error_page 404 /404.html;
            location = /40x.html {
        }

        error_page 500 502 503 504 /50x.html;
            location = /50x.html {
        }

    }
}
```