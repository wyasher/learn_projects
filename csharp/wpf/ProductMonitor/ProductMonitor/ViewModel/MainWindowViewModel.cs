using System;
using System.Collections.Generic;
using System.Collections.Specialized;
using System.ComponentModel;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Controls;
using ProductMonitor.Model;
using ProductMonitor.UserControls;

namespace ProductMonitor.ViewModel
{

    internal class MainWindowViewModel : INotifyPropertyChanged
    {


        public MainWindowViewModel()
        {
            EnvironmentModels.Add(new EnvironmentModel
            {
                ItemTitle = "光照(Lux)",
                ItemValue = 123
            });
            EnvironmentModels.Add(new EnvironmentModel
            {
                ItemTitle = "噪音(db)",
                ItemValue = 55
            });
            EnvironmentModels.Add(new EnvironmentModel
            {
                ItemTitle = "温度(℃)",
                ItemValue = 80
            });
            EnvironmentModels.Add(new EnvironmentModel
            {
                ItemTitle = "湿度(%)",
                ItemValue = 43
            });
            EnvironmentModels.Add(new EnvironmentModel
            {
                ItemTitle = "PM2.5(m3)",
                ItemValue = 123
            });
            EnvironmentModels.Add(new EnvironmentModel
            {
                ItemTitle = "硫化氢(PPM)",
                ItemValue = 123
            });
            EnvironmentModels.Add(new EnvironmentModel
            {
                ItemTitle = "氮气(15)",
                ItemValue = 123
            });
            DeviceModels.Add(new DeviceModel
            {
                ItemTitle = "电压(Kw.h)",
                ItemValue = 11.4,
            });
            DeviceModels.Add(new DeviceModel
            {
                ItemTitle = "电压(V)",
                ItemValue = 222.4,
            });
            DeviceModels.Add(new DeviceModel
            {
                ItemTitle = "电流(A)",
                ItemValue = 5,
            });
            DeviceModels.Add(new DeviceModel
            {
                ItemTitle = "压差(kpa)",
                ItemValue = 13,
            });
            DeviceModels.Add(new DeviceModel
            {
                ItemTitle = "温差(℃)",
                ItemValue = 36,
            });
            DeviceModels.Add(new DeviceModel
            {
                ItemTitle = "震动(mm/s)",
                ItemValue = 4.1,
            });
            DeviceModels.Add(new DeviceModel
            {
                ItemTitle = "转速(r/min)",
                ItemValue = 2600,
            });

            DeviceModels.Add(new DeviceModel
            {
                ItemTitle = "气压(kpa)",
                ItemValue = 0.3,
            });

            AlarmModels.Add(new AlarmModel
            {
                Num = "01",
                Msg = "设备温度过高",
                Time = "2024-11-23 11:11:11",
                Duration = 10
            });
            AlarmModels.Add(new AlarmModel
            {
                Num = "02",
                Msg = "车间温度过高",
                Time = "2024-11-23 11:11:11",
                Duration = 13
            });
            AlarmModels.Add(new AlarmModel
            {
                Num = "03",
                Msg = "设备转速过快",
                Time = "2024-11-23 14:11:11",
                Duration = 13
            });
            AlarmModels.Add(new AlarmModel
            {
                Num = "04",
                Msg = "设备气压过低",
                Time = "2024-11-23 20:11:11",
                Duration = 20
            });

            StuffOutWorkModels =
            [
                new StuffOutWorkModel { StuffName = "张晓婷", Position = "技术员", OutWorkCount = 123 },
                new StuffOutWorkModel { StuffName = "李晓", Position = "操作员", OutWorkCount = 23 },
                new StuffOutWorkModel { StuffName = "王克俭", Position = "技术员", OutWorkCount = 134 },
                new StuffOutWorkModel { StuffName = "陈家栋", Position = "统计员", OutWorkCount = 143 },
                new StuffOutWorkModel { StuffName = "杨过", Position = "技术员", OutWorkCount = 12 },
            ];


            RadarModels.Add(new RadarModel { ItemName = "排烟风机", Value = 90 });
            RadarModels.Add(new RadarModel { ItemName = "客梯", Value = 30.00 });
            RadarModels.Add(new RadarModel { ItemName = "供水机", Value = 34.89 });
            RadarModels.Add(new RadarModel { ItemName = "喷淋水泵", Value = 69.59 });
            RadarModels.Add(new RadarModel { ItemName = "稳压设备", Value = 20 });

            WorkShopModels.Add(new WorkShopModel { WorkShopName = "贴片车间", WorkingCount = 32, WaitCount = 8, WrongCount = 4, StopCount = 0 });
            WorkShopModels.Add(new WorkShopModel { WorkShopName = "封装车间", WorkingCount = 20, WaitCount = 8, WrongCount = 4, StopCount = 0 });
            WorkShopModels.Add(new WorkShopModel { WorkShopName = "焊接车间", WorkingCount = 68, WaitCount = 8, WrongCount = 4, StopCount = 0 });
            WorkShopModels.Add(new WorkShopModel { WorkShopName = "贴片车间", WorkingCount = 68, WaitCount = 8, WrongCount = 4, StopCount = 0 });


            #region 初始化机台列表
            var random = new Random();
            for (var i = 0; i < 20; i++)
            {
                var plan = random.Next(100, 1000);//计划量 随机数
                var finished = random.Next(0, plan);//已完成量
                MachineModels.Add(new MachineModel
                {
                    MachineName = "焊接机-" + (i + 1),
                    FinishedCount = (uint)finished,
                    PlanCount = (uint)plan,
                    Status = "作业中",
                    OrderNo = "H202212345678"
                });
            }
            #endregion


        }

        /// <summary>
        /// 监控用户控件
        /// </summary>
        private UserControl? _monitorUserControl;

        public UserControl MonitorUserControl
        {
            set
            {
                _monitorUserControl = value;
                PropertyChanged?.Invoke(this, new PropertyChangedEventArgs("MonitorUserControl"));
            }
            get
            {
                _monitorUserControl ??= new MonitorUserControl();
                return _monitorUserControl;
            }
        }

        public event PropertyChangedEventHandler? PropertyChanged;

        public string TimeStr => DateTime.Now.ToString("HH:mm");
        public string DateStr => DateTime.Now.ToString("yyyy-MM-dd");

        private string[] weekDays = { "星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期天" };

        public string WeekStr
        {
            get
            {
                int index = (int)DateTime.Now.DayOfWeek;
                return weekDays[index];
            }
        }

        private string _machineCount = "0298";

        public string MachineCount
        {
            set
            {
                _machineCount = value;
               PropertyChanged?.Invoke(this,new PropertyChangedEventArgs("MachineCount"));
            }
            get
            {
                return _machineCount;
            }
        }

        private string _productCount = "1643";

        public string ProductCount
        {
            set
            {
                _productCount = value;
                PropertyChanged?.Invoke(this, new PropertyChangedEventArgs("ProductCount"));
            }
            get
            {
                return _productCount;
            }
        }

        private string _badCount = "1633";

        public string BadCount
        {
            set
            {
                _badCount = value;
                PropertyChanged?.Invoke(this, new PropertyChangedEventArgs("BadCount"));
            }
            get
            {
                return _badCount;
            }
        }

        private List<DeviceModel> _deviceModels = [];

        public List<DeviceModel> DeviceModels
        {
            get => _deviceModels;

            set
            {
                _deviceModels = value;
                PropertyChanged?.Invoke(this, new PropertyChangedEventArgs("DeviceModels"));
            }
        }
        #region 环境监控数据

        private List<EnvironmentModel> _environmentModels = [];

        public List<EnvironmentModel> EnvironmentModels
        {
            get => _environmentModels;

            set
            {
                _environmentModels = value;
                PropertyChanged?.Invoke(this,new PropertyChangedEventArgs("EnvironmentModels"));
            }
        }

        #endregion

        private List<AlarmModel> _alarmModels = [];

        public List<AlarmModel> AlarmModels
        {
            get => _alarmModels;
            set
            {
                _alarmModels = value;
                PropertyChanged?.Invoke(this,new PropertyChangedEventArgs("AlarmModels"));
            }

        }


        private List<RadarModel> _radarModels = [];

        public List<RadarModel> RadarModels
        {
            get => _radarModels;
            set
            {
                _radarModels = value;
                PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(nameof(RadarModels)));

            }
        }

        private List<StuffOutWorkModel> _stuffOutWorkModels = [];

        public List<StuffOutWorkModel> StuffOutWorkModels
        {
            get => _stuffOutWorkModels;
            set
            {
                _stuffOutWorkModels = value;
                PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(nameof(StuffOutWorkModels)));

            }
        }

        private List<WorkShopModel> _workShopModels = [];

        public List<WorkShopModel> WorkShopModels
        {
            get => _workShopModels;
            set
            {
                _workShopModels = value;
                PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(nameof(WorkShopModels)));
            }
        }

        private List<MachineModel> _machineModels = [];

        public List<MachineModel> MachineModels
        {
            get => _machineModels;
            set
            {
                _machineModels = value;
                PropertyChanged?.Invoke(this,new PropertyChangedEventArgs(nameof(MachineModels)));
            }
        }

    }
}
