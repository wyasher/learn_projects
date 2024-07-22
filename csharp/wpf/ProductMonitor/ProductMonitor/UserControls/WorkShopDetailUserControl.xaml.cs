using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
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

namespace ProductMonitor.UserControls
{
    /// <summary>
    /// WorkShopDetailUserControl.xaml 的交互逻辑
    /// </summary>
    public partial class WorkShopDetailUserControl : UserControl
    {
        public WorkShopDetailUserControl()
        {
            InitializeComponent();
        }

        private void OpenDetail(object sender, RoutedEventArgs e)
        {
            DetailMachine.Visibility = Visibility.Visible;
            var thicknessAnimation = new ThicknessAnimation(new Thickness(0,50,0,-50),new Thickness(0,0,0,0),
                
                new TimeSpan(0,0,0,0,400));

            var doubleAnimation = new DoubleAnimation(0, 1, new TimeSpan(0, 0, 0, 0, 400));
            Storyboard.SetTarget(thicknessAnimation,DetailMachineContent);
            Storyboard.SetTarget(doubleAnimation, DetailMachineContent);
            Storyboard.SetTargetProperty(thicknessAnimation,new PropertyPath("Margin"));
            Storyboard.SetTargetProperty(doubleAnimation, new PropertyPath("Opacity"));
            var storyBoard = new Storyboard();
            storyBoard.Children.Add(thicknessAnimation);
            storyBoard.Children.Add(doubleAnimation);
            storyBoard.Begin();
        }

        private void CloseDetail(object sender, RoutedEventArgs e)
        {
            DetailMachine.Visibility = Visibility.Collapsed;
        }
    }

}
