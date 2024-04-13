import FormInfo from '../bean/FormInfo';
import { RDB_NAME } from '../constants/CommonConstants';
import RdbUtils from '../rdb/RdbUtils';

class FormUtils {
  public insertFromData(context: Context, fromInfo: FormInfo): void {
    RdbUtils.initDb(context, RDB_NAME.dbName);
  }
}

export default new FormUtils();