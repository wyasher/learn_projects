export default class DayInfo{
  date: string;
  targetTaskNum: number;
  finTaskNum: number;

  constructor(date: string, targetTaskNum: number,finTaskNum: number) {
    this.date = date;
    this.targetTaskNum = targetTaskNum;
    this.finTaskNum = finTaskNum;
  }
}