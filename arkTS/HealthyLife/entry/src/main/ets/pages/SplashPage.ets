import common from '@ohos.app.ability.common'
import data_preferences from '@ohos.data.preferences'
import Logger from '../common/utils/Logger';

const H_STORE: string = 'healthAppStore';
const IS_PRIVACY: string = 'isPrivacy';
import * as commonConst from '../common/constants/CommonConstants';
import router from '@ohos.router';
import UserPrivacyDialog from '../view/dialog/UserPrivacyDialog';


@Entry
@Component
struct SplashPage {
  context: common.UIAbilityContext = getContext(this) as common.UIAbilityContext;
  dialogController: CustomDialogController = new CustomDialogController({
    builder: UserPrivacyDialog({
      cancel: this.exitApp.bind(this),
      confirm: this.onConfirm.bind(this)
    }),
    cancel: this.exitApp.bind(this),
    autoCancel: false,
    alignment: DialogAlignment.Bottom,
    offset: { dx: 0, dy: commonConst.OFFSET_24 }
  })

  onConfirm() {
    let preferences = data_preferences.getPreferences(this.context, H_STORE);
    preferences.then((res) => {
      res.put(IS_PRIVACY, true).then(() => {
        res.flush();
        Logger.info('SplashPage', 'isPrivacy is put success');
      }).catch((err) => {
        Logger.info('SplashPage', 'isPrivacy put failed. Cause:' + err);
      })
    })
    this.jumpAdPage();
  }

  exitApp() {
    this.context.terminateSelf();
  }

  jumpAdPage() {
    setTimeout(() => {
      router.replaceUrl({ url: "pages/AdvertisingPage" })
    }, commonConst.LAUNCHER_DELAY_TIME)
  }

  aboutToAppear() {
    let preferences = data_preferences.getPreferences(this.context, H_STORE);
    preferences.then((res) => {
      res.get(IS_PRIVACY, false).then((isPrivate) => {
        if (isPrivate) {
          this.jumpAdPage();
        } else {
          this.dialogController.open();
        }
      })
    })
  }
  aboutToDisappear() {
    clearTimeout();
  }
  build() {
    Column() {
      Image($r('app.media.logo'))
        .width($r('app.float.default_120'))
        .aspectRatio(1)
        .margin({ top: $r('app.float.default_120') });
      Text($r('app.string.EntryAbility_label'))
        .fontFamily($r('app.string.HarmonyHeiTi_Bold'))
        .fontSize($r('app.float.default_24'))
        .fontColor($r('app.color.titleColor'))
        .fontWeight(commonConst.FONT_WEIGHT_700)
        .letterSpacing(commonConst.LETTER_1)
        .margin({
          top: $r('app.float.default_20'),
          bottom: $r('app.float.default_8') })
      Text($r('app.string.EntryAbility_desc'))
        .fontFamily($r('app.string.HarmonyHeiTi'))
        .fontColor($r('app.color.titleColor'))
        .fontWeight(commonConst.FONT_WEIGHT_400)
        .letterSpacing(commonConst.LETTER_34)
        .opacity(commonConst.OPACITY_6)
        .fontSize($r('app.float.default_16'))
    }
    .width('100%')
    .height('100%')
    .backgroundImagePosition({x: 0, y: 0})
    .backgroundImage($r('app.media.ic_splash_bg'))
    .backgroundImageSize({ width: '100%', height: '100%' })
  }

}