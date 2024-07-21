using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;

namespace ProductMonitor.UserControls
{
    /// <summary>
    /// RingUserControl.xaml 的交互逻辑
    /// </summary>
    public partial class RingUserControl : UserControl
    {
        public RingUserControl()
        {
            InitializeComponent();

            SizeChanged += OnSizeChanged;
        }

        private void OnSizeChanged(object sender,SizeChangedEventArgs e)
        {
            Draw();
        }


        public double PercentValue
        {
            get => (double)GetValue(PercentValueProperty);

            set => SetValue(PercentValueProperty,value);
        }

        public static readonly DependencyProperty PercentValueProperty = 
            DependencyProperty.Register(nameof(PercentValue),typeof(double),typeof(RingUserControl));

        /// <summary>
        /// 画圆环
        /// </summary>
        private void Draw()
        {
            LayoutOutGrid.Width = Math.Min(RenderSize.Width, RenderSize.Height);
            double raduis = LayoutOutGrid.Width / 2;

            double x = raduis + (raduis - 3) * Math.Cos((PercentValue % 100 * 3.6 - 90) * Math.PI / 180);
            double y = raduis + (raduis - 3) * Math.Sin((PercentValue % 100 * 3.6 - 90) * Math.PI / 180);

            int Is50 = PercentValue < 50 ? 0 : 1;

            //M:移动  A:画弧
            string pathStr = $"M{raduis + 0.01} 3A{raduis - 3} {raduis - 3} 0 {Is50} 1 {x} {y}";//移动路径

            //几何图形对象
            var converter = TypeDescriptor.GetConverter(typeof(Geometry));
            ProcessPath.Data = (Geometry)converter.ConvertFrom(pathStr); ;

        }
    }

    
}
