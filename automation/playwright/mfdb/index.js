// @CreateTime: 2021/6/28
// @Author: key7men
// @Contact: key7men@gmail.com
// @Last Modified By: key7men
// @Last Modified Time: 3:31 PM
// @Description: 当前自动化程序完成两类任务，一类是捕获并处理数据，推送到飞书机器人；另一类是完成自动化发帖工作

const dotenv = require('dotenv');
const { webkit } = require('playwright');

// 读取环境变量，不允许将账户密码信息提供到github上
dotenv.config();
const UMAccount = process.env.UMAccount;
const UMPasswd = process.env.UMPasswd;

// 截图存放地址
const ScreenShotDir = './screenshot';

(async () => {

    // 如果想看到实时操作，可以将headless更改为false
    const browser = await webkit.launch({ headless: false });

    // ------------------------ Task 1: 获取友盟统计信息  ----------------------
    const umPage = await browser.newPage();
    await umPage.goto('https://passport.umeng.com');

    await umPage.frame({
        name: 'alibaba-login-box'
    }).click('[placeholder="【友盟+】账号/CNZZ账号"]');

    await umPage.frame({
        name: 'alibaba-login-box'
    }).fill('[placeholder="【友盟+】账号/CNZZ账号"]', UMAccount);

    await umPage.frame({
        name: 'alibaba-login-box'
    }).click('[placeholder="登录密码"]');

    await umPage.frame({
        name: 'alibaba-login-box'
    }).fill('[placeholder="登录密码"]', UMPasswd);

    await Promise.all([
        umPage.waitForNavigation({ url: 'https://workbench.umeng.com/' }),
        umPage.frame({
            name: 'alibaba-login-box'
        }).click('input:has-text("登录")')
    ]);

    // Click li:nth-child(2) .img
    const [umPage1] = await Promise.all([
        umPage.waitForEvent('popup'),
        umPage.click('li:nth-child(2) .img')
    ]);

    // Click text=立即使用
    await umPage1.click('text=立即使用');

    // Click :nth-match(:text("查看报表"), 3)
    const [umPageDetail] = await Promise.all([
        umPage1.waitForEvent('popup'),
        umPage1.click(':nth-match(:text("查看报表"), 3)'),
    ]);

    await umPageDetail.click('text=网站概况');
    await umPageDetail.click('text=历史累计'); // 当历时累计出现才会认定加载成功
    const statisticWin = await umPageDetail.$('#overview_top_order_table');
    const picName = new Date().getTime() + '.png';
    await statisticWin.screenshot({ path: ScreenShotDir + '/' + picName });

    // ------------------------ Task2: 获取Memfiredb用户数据 ----------------------
    const context = await browser.newContext({ storageState: 'mfdb.auth.json' })
    const mfPage = await context.newPage()
    await mfPage.goto('https://cloud.memfiredb.com:9443');
    const totalUsers = await mfPage.innerText('li.ant-pagination-total-text');
    console.log(totalUsers);
    await browser.close();
})();
