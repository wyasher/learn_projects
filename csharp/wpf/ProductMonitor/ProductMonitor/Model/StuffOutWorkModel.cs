using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ProductMonitor.Model
{
    public class StuffOutWorkModel
    {
        public string StuffName { get; set; }
        public string Position { get; set;}
        public uint OutWorkCount { get; set;}


        public uint ShowWidth => OutWorkCount * 10 / 100;
    }
}
