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

        private string[] weekDays = {"星期一","星期二","星期三","星期四","星期五","星期六","星期天" };

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
    }
}
