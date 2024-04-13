import DayInfo from '../bean/DayInfo';
import dataRdb from '@ohos.data.relationalStore';
import { DAY_INFO } from '../constants/CommonConstants';
import RdbUtils from '../rdb/RdbUtils';
import Logger from '../utils/Logger';

class DayInfoApi {
  insertData(dayInfo: DayInfo, callback: Function): void {
    const valueBucket = generateBucket(dayInfo);
    RdbUtils.insert(DAY_INFO.tableName, valueBucket)
      .then(result => callback(result));
    Logger.info('DayInfoTable', 'Insert dayInfo finished.')
  }

  updateData(dayInfo: DayInfo, callback: Function): void {
    const valueBucket = generateBucket(dayInfo);
    let predicates = new dataRdb.RdbPredicates(DAY_INFO.tableName);
    predicates.equalTo("date", dayInfo.date);
    RdbUtils.update(valueBucket, predicates).then(result => callback(result));
    Logger.info('DayInfoTable', 'Update dayInfo finished.');
  }
}

function generateBucket(dayInfo: DayInfo): dataRdb.ValuesBucket {
  let valueBucket = {};
  DAY_INFO.columns.forEach((item: string) => {
    valueBucket[item] = dayInfo[item];
  });
  return valueBucket;
}

let dayInfoApi = new DayInfoApi();

export default dayInfoApi as DayInfoApi;