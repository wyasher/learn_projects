using ProductMonitor.Model;
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
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;

namespace ProductMonitor.UserControls
{
    /// <summary>
    /// RadarUserControl.xaml 的交互逻辑
    /// </summary>
    public partial class RadarUserControl : UserControl
    {
        public RadarUserControl()
        {
            InitializeComponent();
            SizeChanged += OnSizeChanged;
        }

        private void OnSizeChanged(object sender, SizeChangedEventArgs e)
        {
            Drag();
        }

        /// <summary>
        /// 画图
        /// </summary>
        private void Drag()
        {
            if (ItemSource == null || ItemSource.Count == 0)
            {
                return;
            }
            // 清空之前的画布
            MainCanvas.Children.Clear();
            P1.Points.Clear();
            P2.Points.Clear();
            P3.Points.Clear();
            P4.Points.Clear();
            P5.Points.Clear();
            // 调整大小
            var size = Math.Min(RenderSize.Width, RenderSize.Height);
            LayoutGrid.Height = size;
            LayoutGrid.Width = size;
            // 半径
            var radius = size / 2;
            var step = 360.0 / ItemSource.Count;

            for (var i = 0; i < ItemSource.Count; i++)
            {
                var x = (radius - 20) * Math.Cos((step * i - 90) * Math.PI / 180);//x偏移量
                var y = (radius - 20) * Math.Sin((step * i - 90) * Math.PI / 180);//y偏移量

                //X Y坐标
                P1.Points.Add(new Point(radius + x, radius + y));

                P2.Points.Add(new Point(radius + x * 0.75, radius + y * 0.75));

                P3.Points.Add(new Point(radius + x * 0.5, radius + y * 0.5));

                P4.Points.Add(new Point(radius + x * 0.25, radius + y * 0.25));

                //数据多边形
                P5.Points.Add(new Point(radius + x * ItemSource[i].Value * 0.01, radius + y * ItemSource[i].Value * 0.01));

                //文字处理
                TextBlock txt = new()
                {
                    Width = 60,
                    FontSize = 10,
                    TextAlignment = TextAlignment.Center,
                    Text = ItemSource[i].ItemName,
                    Foreground = new SolidColorBrush(Color.FromArgb(100, 255, 255, 255))
                };
                txt.SetValue(Canvas.LeftProperty, radius + (radius - 10) * Math.Cos((step * i - 90) * Math.PI / 180) - 30);//设置左边间距
                txt.SetValue(Canvas.TopProperty, radius + (radius - 10) * Math.Sin((step * i - 90) * Math.PI / 180) - 7);//设置上边间距

                MainCanvas.Children.Add(txt);

            }


        }

        /// <summary>
        /// 依赖属性
        /// </summary>

        public List<RadarModel> ItemSource
        {
            get => (List<RadarModel>)GetValue(ItemSourceProperty);
            set
            {
                SetValue(ItemSourceProperty,value);
            }
        }

        public static readonly DependencyProperty ItemSourceProperty = DependencyProperty.Register(nameof(ItemSource), typeof(List<RadarModel>), typeof(RadarUserControl));

    }
}
