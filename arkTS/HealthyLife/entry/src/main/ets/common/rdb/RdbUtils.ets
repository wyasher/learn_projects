import Logger from '../utils/Logger';
import { RdbHelper } from './RdbHelper';
import dataRdb from '@ohos.data.relationalStore';
import {RdbHelperImpl} from './RdbHelperImpl';
import ColumnInfo from '../bean/ColumnInfo';
let dbContext: Context;
let mDatabaseName: string = '';
export class  RdbUtils{
  private rdbHelpers = new Map<string, RdbHelper>();

  initDb(context: Context, databaseName: string) {
    dbContext = context;
    mDatabaseName = databaseName;
  }

  createDb(): Promise<RdbHelper> {
    return new Promise<RdbHelper>((success, error) => {
      let dbName = mDatabaseName;
      if (!dbContext || !dbName || dbName.length === 0) {
        error("init err");
        return;
      }
      let dbHelper = this.rdbHelpers.get(dbName);
      if (!dbHelper) {
        Logger.info(`initRdb  RdbUtils success`);
        let rdbHelper: RdbHelper = new RdbHelperImpl(dbName);
        rdbHelper.getRdb(dbContext).then(data => {
          this.rdbHelpers.set(dbName, data);
          success(data);
        }).catch(err => {
          error(err);
        })
      } else {
        success(dbHelper);
      }
    })
  }

  deleteDb(context: any, dbName: string): Promise<void> {
    this.rdbHelpers.delete(dbName);
    return dataRdb.deleteRdbStore(context, dbName);
  }

  createTable(tableName: string, columns: Array<ColumnInfo>): Promise<void> {
    return this.createDb().then(dbHelper => {
      return dbHelper.createTable(tableName, columns);
    })
  }

  isCreateTable(tableName: string, columns: Array<ColumnInfo>): Promise<boolean> {
    return this.createTable(tableName, columns).then(() => {
      return true;
    }).catch((error) => {
      Logger.error('RdbUtils', 'create table error ' + JSON.stringify(error));
      return false;
    });
  }

  deleteTable(tableName: string): Promise<void> {
    return this.createDb().then(dbHelper => {
      return dbHelper.deleteTable(tableName);
    })
  }

  executeSql(sql: string): Promise<void> {
    return this.createDb().then(dbHelper => {
      return dbHelper.executeSql(sql);
    })
  }

  addTableColumn(tableName: string, column: ColumnInfo): Promise<void> {
    return this.createDb().then(dbHelper => {
      return dbHelper.addTableColumn(tableName, column);
    })
  }

  insert(tableName: string, values: dataRdb.ValuesBucket | Array<dataRdb.ValuesBucket>): Promise<number> {
    return this.createDb().then(dbHelper => {
      return dbHelper.insert(tableName, values);
    })
  }

  update(values: dataRdb.ValuesBucket, rdbPredicates: dataRdb.RdbPredicates): Promise<number> {
    return this.createDb().then(dbHelper => {
      return dbHelper.update(values, rdbPredicates);
    })
  }

  query(rdbPredicates: dataRdb.RdbPredicates, columns?: Array<string>): Promise<dataRdb.ResultSet> {
    return this.createDb().then(dbHelper => {
      return dbHelper.query(rdbPredicates, columns);
    })
  }

  queryAll(tableName: string): Promise<dataRdb.ResultSet> {
    return this.createDb().then(dbHelper => {
      return dbHelper.queryAll(tableName);
    })
  }

  queryBySql(sql: string, bindArgs?: Array<dataRdb.ValueType>): Promise<dataRdb.ResultSet> {
    return this.createDb().then(dbHelper => {
      return dbHelper.queryBySql(sql, bindArgs);
    })
  }

  del(rdbPredicates: dataRdb.RdbPredicates): Promise<number> {
    return this.createDb().then(dbHelper => {
      return dbHelper.delete(rdbPredicates);
    })
  }
}

let rdbUtils = new RdbUtils();

export default rdbUtils as RdbUtils;