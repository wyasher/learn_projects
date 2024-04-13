import Logger from './Logger';

const FILE_TAG = 'BroadCast';

export class Broadcast {
  private callBackArray = {};

  public on(event, callback) {
    Logger.info(FILE_TAG, 'register broadcast with type '+ event);
    (this.callBackArray[event] || (this.callBackArray[event] = [])).push(callback);
  }

  public off(event, callback) {
    if (event == null) {
      Logger.info(FILE_TAG, 'cancel all broadcast');
      this.callBackArray = {};
    }
    Logger.info(FILE_TAG, 'cancel broadcast with type '+ event);
    const cbs = this.callBackArray[event];
    if (!cbs) {
      return;
    }
    if (callback == null) {
      this.callBackArray[event] = null;
    }

    cbs.splice(cbs.indexOf(callback), 1);
  }

  public emit(event, args?: any[]) {
    if (!this.callBackArray[event]) {
      Logger.info(FILE_TAG, 'emit broadcast failed for no callback');
      return;
    }
    Logger.info(FILE_TAG, 'emit broadcast with type '+ event);
    let cbs = [...this.callBackArray[event]];
    if (cbs) {
      let len = cbs.length;
      for (let i = 0; i < len; i++) {
        try {
          cbs[i].apply(this, args);
        } catch (error) {
          new Error(error);
        }
      }
    }
  }
}

export enum BoardCastType{
  SHOW_ACHIEVEMENT_DIALOG = 'showAchievementDialog',
  SHOW_TASK_DETAIL_DIALOG = 'showTaskDetailDialog',
  SHOW_TARGET_SETTING_DIALOG = 'showTargetSettingDialog',
  SHOW_REMIND_TIME_DIALOG = 'showRemindTimeDialog',
  SHOW_FREQUENCY_DIALOG = 'showFrequencyDialog'
}