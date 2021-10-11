<?php
// var_dump($_SERVER);die();
// define( 'WP_MEMORY_LIMIT', '96M' );
$http_url='http://'.$_SERVER['HTTP_HOST'];

define('WP_HOME', $http_url);
define('WP_SITEURL', $http_url);
define('WPLANG', 'zh_CN');
// define('WP_POST_REVISIONS', false); //是否启用版本修订功能
define('WP_POST_REVISIONS', 3); //最大修订数功能

// define( 'CUSTOM_USER_TABLE', $table_prefix.'my_users' );
// define( 'CUSTOM_USER_META_TABLE', $table_prefix.'my_usermeta' );
// define( 'SAVEQUERIES', true );
$_no_output_and_use_but_SAVEQUERIES_true = <<<EOT
<?php
if ( current_user_can( 'administrator' ) ) {
    global $wpdb;
    echo "<pre>";
    print_r( $wpdb->queries );
    echo "</pre>";
}
?>
EOT;
// define( 'FS_CHMOD_DIR', ( 0755 & ~ umask() ) );
// define( 'FS_CHMOD_FILE', ( 0644 & ~ umask() ) );
define( 'DISABLE_WP_CRON', true );
// define( 'WP_ALLOW_REPAIR', true ); // 当无法登录，无法访问数据库时，通过此项来修复数据库，仅需要用时开启此功能
define( 'DISALLOW_FILE_EDIT', true );
// define( 'FORCE_SSL_ADMIN', true ); // 后台必须通过加密才能链接
define( 'WP_HTTP_BLOCK_EXTERNAL', true );
define( 'WP_ACCESSIBLE_HOSTS', 'api.wordpress.org,*.github.com' );
define( 'AUTOMATIC_UPDATER_DISABLED', true );
define( 'WP_AUTO_UPDATE_CORE', false ); // 关闭核心更新
define( 'DISALLOW_UNFILTERED_HTML', true ); //不知其它用户要被禁止html数据，超级管理员也必须禁止
/**
 * WordPress基础配置文件。
 *
 * 这个文件被安装程序用于自动生成wp-config.php配置文件，
 * 您可以不使用网站，您需要手动复制这个文件，
 * 并重命名为“wp-config.php”，然后填入相关信息。
 *
 * 本文件包含以下配置选项：
 *
 * * MySQL设置
 * * 密钥
 * * 数据库表名前缀
 * * ABSPATH
 *
 * @link https://codex.wordpress.org/zh-cn:%E7%BC%96%E8%BE%91_wp-config.php
 *
 * @package WordPress
 */

// ** MySQL 设置 - 具体信息来自您正在使用的主机 ** //
/** WordPress数据库的名称 */
define('DB_NAME', 'wordpress');

/** MySQL数据库用户名 */
define('DB_USER', 'wordpress');

/** MySQL数据库密码 */
define('DB_PASSWORD', 'wordpress');

/** MySQL主机 */
define('DB_HOST', 'db:3306');

/** 创建数据表时默认的文字编码 */
define('DB_CHARSET', 'utf8');

/** 数据库整理类型。如不确定请勿更改 */
define('DB_COLLATE', '');

/**#@+
 * 身份认证密钥与盐。
 *
 * 修改为任意独一无二的字串！
 * 或者直接访问{@link https://api.wordpress.org/secret-key/1.1/salt/
 * WordPress.org密钥生成服务}
 * 任何修改都会导致所有cookies失效，所有用户将必须重新登录。
 *
 * @since 2.6.0
 */
define('AUTH_KEY',         'SddE[*nTCqbUopu4hX2t<@Y;~#wCxgv6~$_?)FZI<%R=S4|,~1bzwQ{m#Os)l!@K');
define('SECURE_AUTH_KEY',  '?aT*9/Oc-IrEa8C]lgX01xJ*ZRdjWA&3hI#}U5fj-hfI3^mXPs=Fo7I(46OV&q_/');
define('LOGGED_IN_KEY',    'N0[)9;_rLVO`G:oF+JQ=+mE+zq#I[]#2GTmj<n}RJ!UXL<bB2@:9(c8N?JgF6M-D');
define('NONCE_KEY',        '*^Qg|r:/t :@A6X:jHpYGF5y*Z5Ae!L17?JW0#!M^+$F+}+hu~FO3]sNuP#h{Hlr');
define('AUTH_SALT',        '2Q.4lXWZeBMs6~ )a5v,TN[A@ZUx/,9awX,[x(cV^`Y(FC{]}vmBFZz{YJ[vc n4');
define('SECURE_AUTH_SALT', '<wDk&WHymh/_YA2v7zW.QV3f]b=Zae&oK2i 5eiktU&JVMTy[Wa2U^0.<SW;J@[,');
define('LOGGED_IN_SALT',   'e!n0$`V#>|gxEivz#[q4N/LZ1(](bC@Z37pc_`<@mHi8W~QpNZ.B]H<|UvwV.Rj{');
define('NONCE_SALT',       ']SGFsbd2IIM{|Bs|1]C6$b$vHDum:{qLXyNQ+x[H^Xx}pbz)ljWa#w2c7EN_e-{D');
/**#@-*/

/**
 * WordPress数据表前缀。
 *
 * 如果您有在同一数据库内安装多个WordPress的需求，请为每个WordPress设置
 * 不同的数据表前缀。前缀名只能为数字、字母加下划线。
 */
$table_prefix  = 'www_';

/**
 * 开发者专用：WordPress调试模式。
 *
 * 将这个值改为true，WordPress将显示所有用于开发的提示。
 * 强烈建议插件开发者在开发环境中启用WP_DEBUG。
 *
 * 要获取其他能用于调试的信息，请访问Codex。
 *
 * @link https://codex.wordpress.org/Debugging_in_WordPress
 */
defined('WP_DEBUG') || define('WP_DEBUG', true); //兼容WP—CLI
if ( WP_DEBUG ) {
    defined('WP_DEBUG_DISPLAY') || define('WP_DEBUG_DISPLAY', false);
    defined('WP_DEBUG_LOG') || define('WP_DEBUG_LOG', true);
}

/**
 * zh_CN本地化设置：启用ICP备案号显示
 *
 * 可在设置→常规中修改。
 * 如需禁用，请移除或注释掉本行。
 */
define('WP_ZH_CN_ICP_NUM', true);
if (isset($_SERVER['HTTP_X_FORWARDED_PROTO']) && $_SERVER['HTTP_X_FORWARDED_PROTO'] === 'https') {
	$_SERVER['HTTPS'] = 'on';
}
/* 好了！请不要再继续编辑。请保存本文件。使用愉快！ */

/** WordPress目录的绝对路径。 */
if ( !defined('ABSPATH') )
	define('ABSPATH', dirname(__FILE__) . '/');

/** 设置WordPress变量和包含文件。 */
require_once(ABSPATH . 'wp-settings.php');