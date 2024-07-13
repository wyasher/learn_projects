using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ProductMonitor.Model
{
    public class WorkShopModel
    {
        public string WorkShopName { get; set; }
        public uint WorkingCount { get; set; }
        public uint WaitCount { get; set; }
        public uint WrongCount { get; set; }
        public uint StopCount { get; set; }

        public uint TotalCount
        {
            get => WaitCount + WrongCount + StopCount;

    }
    }
}
