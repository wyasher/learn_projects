using System.Text;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Animation;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;
using ProductMonitor.UserControls;
using ProductMonitor.ViewModel;

namespace ProductMonitor
{
    /// <summary>
    /// Interaction logic for MainWindow.xaml
    /// </summary>
    public partial class MainWindow : Window
    {
        private MainWindowViewModel mainWindowViewModel = new MainWindowViewModel();
        public MainWindow()
        {
            InitializeComponent();
             mainWindowViewModel = new MainWindowViewModel();
            DataContext = mainWindowViewModel;
        }
        /// <summary>
        ///
        /// 显示车间详情页
        /// </summary>
        private void ShowDetailUserControl()
        {
            var workShopDetail = new WorkShopDetailUserControl();
            mainWindowViewModel.MonitorUserControl = workShopDetail;
            // 动画效果
            // 位移和移动的时间
            var thicknessAnimation = new ThicknessAnimation(new Thickness(0,50,0,-50),new Thickness(0,0,0,0),new Duration(new TimeSpan(0,0,0, 0, 200)));
            // 透明度
            var doubleAnimation = new DoubleAnimation(0, 1, new Duration(new TimeSpan(0, 0, 0, 0,200)));
            Storyboard.SetTarget(thicknessAnimation,workShopDetail);
            Storyboard.SetTarget(doubleAnimation, workShopDetail);

            Storyboard.SetTargetProperty(thicknessAnimation,new PropertyPath("Margin"));
            Storyboard.SetTargetProperty(doubleAnimation, new PropertyPath("Opacity"));
            
            var storyboard = new Storyboard();
            storyboard.Children.Add(thicknessAnimation);
            storyboard.Children.Add(doubleAnimation);
            storyboard.Begin();

        }

        public Command.Command ShowDetailCommand => new(ShowDetailUserControl);

        public void GoBackMonitor()
        {
            var monitor
                = new MonitorUserControl();
            mainWindowViewModel.MonitorUserControl = monitor;
        }



        public Command.Command GoBackMonitorCommand => new(GoBackMonitor);

        // 最小化
        private void BtnMin(object sender,RoutedEventArgs e)
        {
            WindowState = WindowState.Minimized;
        }
        
        // 关闭

        private void BtnClose(object sender, RoutedEventArgs e)
        {
            Environment.Exit(0);
        }

        private void ShowSettingWindow()
        {
            var settingsWindow = new SettingsWindow()
            {
                Owner = this
            };
            settingsWindow.ShowDialog();
        }

        public Command.Command ShowSettingCommand => new(ShowSettingWindow);
    }
  
}