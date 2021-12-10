// Fisplus 模块配置
// fisp release -wd test100


// 基本配置
fis.config.merge({
    namespace: 'phplib'
});

// 部署配置。
var deployTargets = {
    test114: {
        host: '192.168.240.137',
        port: 8020
    }
};

var deployConfig = [{
    from: '/ds',
    to: '/home/homework/php/phplib/kunpeng/ds',
    subOnly: true
},{
    from: '/util',
    to: '/home/homework/php/phplib/kunpeng/util',
    subOnly: true
},{
    from: '/conf',
    to: '/home/homework/php/phplib/kunpeng/conf',
    subOnly: true
},{
    from: '/dao',
    to: '/home/homework/php/phplib/kunpeng/dao',
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