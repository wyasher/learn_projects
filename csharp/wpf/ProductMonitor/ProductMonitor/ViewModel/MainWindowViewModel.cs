using System;
using System.Collections.Generic;
using System.Collections.Specialized;
using System.ComponentModel;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Controls;
using ProductMonitor.UserControls;

namespace ProductMonitor.ViewModel
{

    internal class MainWindowViewModel : INotifyPropertyChanged
    {
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
    }
}
