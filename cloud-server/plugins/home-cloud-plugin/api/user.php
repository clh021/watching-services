<?php
function homecloud_user_register_handler($request = null)
{
    $response = array();
    $body_params = $request->get_body_params();
    $username = $body_params['username'];
    $password = $body_params['password'];
    $email = $body_params['email'];
    // $role = $body_params['role'];
    $error = new WP_Error();
    if (empty($username)) {
        $error->add(400, __("Username field 'username' is required.", 'wp-rest-user'), array('status' => 400));
        return $error;
    }
    if (empty($email)) {
        $error->add(401, __("Email field 'email' is required.", 'wp-rest-user'), array('status' => 400));
        return $error;
    }
    if (empty($password)) {
        $error->add(404, __("Password field 'password' is required.", 'wp-rest-user'), array('status' => 400));
        return $error;
    }
    $user_id = username_exists($username);
    if (!$user_id && email_exists($email) == false) {
        $user_id = wp_create_user($username, $password, $email);
        if (!is_wp_error($user_id)) {
            // Ger User Meta Data (Sensitive, Password included. DO NOT pass to front end.)
            $user = get_user_by('id', $user_id);
            // $user->set_role($role);
            $user->set_role('subscriber');
            // WooCommerce specific code
            if (class_exists('WooCommerce')) {
                $user->set_role('customer');
            }
            // Ger User Data (Non-Sensitive, Pass to front end.)
            $response['code'] = 200;
            $response['message'] = __("User '" . $username . "' Registration was Successful", "wp-rest-user");
        } else {
            return $user_id;
        }
    } else {
        $error->add(406, __("Email already exists, please try 'Reset Password'", 'wp-rest-user'), array('status' => 400));
        return $error;
    }
    return new WP_REST_Response($response, 123);
}
function homecloud_user_login_handler($request = null)
{
    // $parameters = $request->get_json_params();
    // $body = $request->get_body();
    $body_params = $request->get_body_params();
    $username = $body_params['username'];
    $password = $body_params['password'];

    $data['user_login'] = $username;
    $data['user_password'] = $password;
    $data['remember'] = false;
    $user = wp_signon($data, false);
    // if (!is_wp_error($user)) {
    //     return new WP_REST_Response($user,404);
    // } else {
    //     return new WP_REST_Response($user, 404);
    // }
    // $response['login_user'] = $user;

    // $response['current_user_id'] = get_current_user_id();
    // $response['current_user'] = wp_get_current_user();
    $response['code'] = is_wp_error($user) ? 404 : 200;
    $response['authorization'] = is_wp_error($user) ? "" : base64_encode( $username . ':' . $password );
    $resultMsg = is_wp_error($user) ? 'Login Failed' : 'Login Successful';
    $response['message'] = __("User '" . $username . "' $resultMsg", "wp-rest-user");
    return new WP_REST_Response($response, 123);
}
