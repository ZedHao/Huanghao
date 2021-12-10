// Fisplus 模块配置
// fisp release -wd test100


// 基本配置
fis.config.merge({
    namespace: 'kunpeng'
});

// 部署配置。
var deployTargets = {
    test100: {
        host: '192.168.240.125',
        port: 8020
    }
};

var deployConfig = [{
    from: '/actions',
    to: '/home/homework/app/kunpeng/actions',
    subOnly: true
}, {
    from: '/controllers',
    to: '/home/homework/app/kunpeng/controllers',
    subOnly: true
}, {
    from: '/library',
    to: '/home/homework/app/kunpeng/library',
    subOnly: true
}, {
    from: '/models',
    to: '/home/homework/app/kunpeng/models',
    subOnly: true
}, {
    from: '/script',
    to: '/home/homework/app/kunpeng/script',
    subOnly: true
}, {
    from: '/template',
    to: '/home/homework/template',
    subOnly: true
}, {
    from: '/conf',
    to: '/home/homework/conf/app/kunpeng',
    subOnly: true
}
];

fis.util.map(deployTargets, function (serverName, serverConfig) {
    var _deployConfig = serverConfig.deploy || [].concat(deployConfig);
    for (var i = 0; i < _deployConfig.length; ++i) {
        fis.util.merge(_deployConfig[i], {
            receiver: 'http://' + serverConfig.host + ':' + (serverConfig.port || '8080') + '/fisreceiver.php',
            exclude: /\/page\/demo\//
        });
    }
    fis.config.set('deploy.' + serverName, _deployConfig);
});