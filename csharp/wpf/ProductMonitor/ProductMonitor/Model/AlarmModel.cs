﻿using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ProductMonitor.Model
{
    internal class AlarmModel
    {
        public string Num { get; set; }
        public string Msg { get; set; }
        public string Time { get; set; }
        public uint Duration { get; set; }
    }
}
