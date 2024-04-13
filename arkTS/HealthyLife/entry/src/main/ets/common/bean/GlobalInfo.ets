export default class GlobalInfo {
  firstDate: string;
  lastDate: string;
  checkInDays: string;
  achievements: string;

  constructor(firstDate: string, lastDate: string, checkInDays: string, achievements: string) {
    this.firstDate = firstDate;
    this.lastDate = lastDate;
    this.checkInDays = checkInDays;
    this.achievements = achievements;
  }
}