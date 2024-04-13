import hilog from '@ohos.hilog';

const LOGGER_PREFIX = 'HealthyLife';

class Logger {
  private domain: number;
  private prefix: string;
  private format: string = '%{public}s, %{public}s';

  constructor(prefix: string = '', domain: number = 0xFF00) {
    this.prefix = prefix;
    this.domain = domain;
  }

  debug(...args: any[]): void {
    hilog.debug(this.domain, this.prefix, this.format, ...args);
  }
  info(...args: any[]): void {
    hilog.info(this.domain, this.prefix, this.format, args);
  }

  warn(...args: any[]): void {
    hilog.warn(this.domain, this.prefix, this.format, args);
  }

  error(...args: any[]): void {
    hilog.error(this.domain, this.prefix, this.format, args);
  }
}

export default new Logger(LOGGER_PREFIX,0xFF02)