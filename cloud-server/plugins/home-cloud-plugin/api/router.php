<?php
// 路由
add_action('rest_api_init', function () {

    // 注册 /homecloud/users/register
    // register_rest_route('homecloud', 'users/register', [
    //     'methods' => 'POST',
    //     'callback' => 'homecloud_user_register_handler',
    //     'permission_callback' => __return_true,
    //     'args' => array(
    //         'username' => array(
    //             'required' => true,
    //         ),
    //         'password' => array(
    //             'required' => true,
    //         ),
    //         'email' => array(
    //             'required' => true,
    //         ),
    //     ),
    // ]);
    // 用户注册是所有流程的第一步，设备连接(cloudReg)的前提
    // 注册至少包含三个字段： username, password, email

    // 登陆 /homecloud/users/login
    register_rest_route('homecloud', 'users/login', [
        'methods' => 'POST',
        'callback' => 'homecloud_user_login_handler',
        'permission_callback' => __return_true,
        'args' => array(
            'username' => array(
                'required' => true,
            ),
            'password' => array(
                'required' => true,
            ),
        ),
    ]);
    // 使用 username,password 登陆，拿到 nonce 登陆 key 作为其他认证接口的凭证

    // 设备连接 /homecloud/device/connect
    register_rest_route('homecloud', 'device/connect', [
        'methods' => 'POST',
        'callback' => 'homecloud_device_connect_handler',
        'permission_callback' => function () {
            return is_user_logged_in();
        },
        'args' => array(
            'id' => array('required' => true),
            // 'uptime' => array('required' => true),
            // 'disk' => array('required' => true),
            // 'cpu' => array('required' => true),
            // 'mem' => array('required' => true),
        ),
    ]);
    // 使用 username,password 登陆，拿到 nonce 登陆 key 作为心跳接口的凭证 没有返回 nonce 时会返回 errorMsg 错误

    // 设备状态更新(心跳) /homecloud/device/status
    // register_rest_route('homecloud', 'device/status', [
    //     'methods' => 'POST',
    //     'callback' => 'homecloud_device_status_handler',
    //     'permission_callback' => function () {
    //         return is_user_logged_in();
    //     },
    //     'args' => array(
    //         'status' => array(
    //             'required' => true,
    //         ),
    //     ),
    // ]);
    // 添加 nonce 到请求头，返回 result:true|false 添加安装任务的状态(可能 nonce 会过期, 过期需要重新登陆)
    // 该接口是设备状态列表的数据支撑

    // 设备状态列表(一个用户可能有多个设备)  /homecloud/device/list
    // register_rest_route('homecloud', 'device/list', [
    //     'methods' => 'GET',
    //     'callback' => 'homecloud_device_list_handler',
    //     'permission_callback' => function () {
    //         return is_user_logged_in();
    //     },
    // ]);
    // 添加 nonce 到请求头，返回 deviceList []device
    // 返回：设备列表及设备在线状态(能否连接的因素：帐号登陆|关机了)

    // 应用列表  /homecloud/app/list
    register_rest_route('homecloud', 'app/list', [
        'methods' => 'GET',
        'callback' => 'homecloud_app_list_handler',
        'permission_callback' => function () {
            return is_user_logged_in();
        },
        'args' => array(
            'deviceId' => array(
                'required' => true,
                'validate_callback' => function ($param, $request, $key) {
                    // return is_deviceId_validate($param);
                    return true;
                },
            ),
        ),
    ]);
    // 添加 nonce 到请求头，返回 appList []app
    // 包含 appId 已进行 install 操作

    // 应用安装  /homecloud/app/install
    // register_rest_route('homecloud', 'app/install', [
    //     'methods' => 'POST',
    //     'callback' => 'homecloud_app_install_handler',
    //     'permission_callback' => function () {
    //         return is_user_logged_in();
    //     },
    // ]);
    // 添加 nonce 到请求头，返回 result:true|false 添加安装任务的状态(应用并非立即就能知道是否安装成功)

    // 设备状态  /homecloud/device/info/${DeviceID}
    // register_rest_route('homecloud', 'device/info/${DeviceID}', [
    //     'methods' => 'GET',
    //     'callback' => 'homecloud_device_info_handler',
    //     'permission_callback' => function () {
    //         return is_user_logged_in();
    //     },
    // ]);
    // 添加 nonce 到请求头，设备ID，返回 {deviceInfo, []Task}
    // 返回：设备当前基础信息(磁盘空间，开机时间)， 正在进行的任务列表(安装|卸载 app 的队列)
});
