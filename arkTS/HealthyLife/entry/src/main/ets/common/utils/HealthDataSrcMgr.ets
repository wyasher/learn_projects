import { Broadcast } from './BoradCast';
import Logger from './Logger';

const APP_KEY_GROUP_DATA_SOURCE_MANAGER = 'app_key_group_data_source_manager';

export class HealthDataSrcMgr {
  private broadcast: Broadcast

  constructor() {
    Logger.debug('HealthDataSourceManager', 'constructor');
    this.broadcast = new Broadcast();
  }

  public static getInstance(): HealthDataSrcMgr {
    let healthDataSrcMgr = AppStorage.Get<HealthDataSrcMgr>(APP_KEY_GROUP_DATA_SOURCE_MANAGER)
    if (healthDataSrcMgr == null) {
      healthDataSrcMgr = new HealthDataSrcMgr();
      AppStorage.SetOrCreate<HealthDataSrcMgr>(APP_KEY_GROUP_DATA_SOURCE_MANAGER, healthDataSrcMgr);
    }
    return healthDataSrcMgr;
  }

  public getBordCast(): Broadcast {
    return this.broadcast;
  }
}