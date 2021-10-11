<?php
function homecloud_device_connect_handler() {
    // 获取客户端设备传递的 id
    // 记录：客户端的IP,连接时间
    $body_params = $request->get_body_params();
    $deviceId = $body_params['deviceId'];

$response['code'] = is_wp_error($user) ? 404 : 200;
$response['authorization'] = is_wp_error($user) ? "" : base64_encode($username . ':' . $password);
$resultMsg = is_wp_error($user) ? 'Login Failed' : 'Login Successful';
$response['message'] = __("User '" . $username . "' $resultMsg", "wp-rest-user");
return new WP_REST_Response($response, 123);

}