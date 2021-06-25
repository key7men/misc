// @CreateTime: 2021/6/18
// @Author: key7men
// @Contact: key7men@gmail.com
// @Last Modified By: key7men
// @Last Modified Time: 2:44 PM
// @Description: 每天自动化定时在各个社区动态发状态
// @Target：掘金社区

// Device Info
// width x height = 1080 x 2340
// screen size = 6.53 inch

// 获取当前脚本的路径
let path = engines.myEngine().source.toString();
// 获取当前时间后10分钟的时间戳
let millis = new Date().getTime() + 8 * 60 * 60 * 1000;
// 预定一个10分钟后的任务，这样10分钟后会再次执行本脚本，并再次预定定时任务，从而每10分钟循环
toastLog("掘金定时任务预定成功: ", timers.addDisposableTask({
    path: path,
    date: millis
}));

// 执行主脚本逻辑
main();

function main() {
    // 必须启动手机的无障碍模式
    auto.waitFor();

    // ------------------------ 掘金沸点 ------------------------

    // 通过：adb shell dumpsys window | grep mCurrentFocus
    // 提前查到了包名称与activity名称：com.daimajia.gold/im.juejin.android.ui.MainActivity

    // step0: 获取v2ex上的内容
    const v2exApi = 'https://www.v2ex.com/api/topics/hot.json';
    const res = http.get(v2exApi);
    let content = '';
    if (res.statusCode === 200) {
        const data = JSON.parse(res.body.string());
        content = data[0].title;
    } else {
        toast('请求失败: ' + res.statusMessage);
    }

    // step1: 打开掘金引用，可通过adb获取包名，比应用名称可靠:
    const juejin = 'com.daimajia.gold';
    const juejinActivity = 'im.juejin.android.ui.MainActivity';

    try {
        app.launchPackage(juejin);
    } catch (error) {
        console.error('无法找到掘金应用');
    }

    waitForActivity(juejinActivity, 1000);

    const FeidianBtn = desc('沸点').findOne();
    FeidianBtn.click();

    const PostBtn = id('btn_post').findOne();
    PostBtn.click();

    const PostInput = id('et_pin').findOne();
    PostInput.setText('一日三省，灵魂一问：' + content + '?');

    const PublishBtn = id('action_post').findOne();
    PublishBtn.click();

    app.openAppSetting(juejin);
    while(!click('强制停止'));
}











































































