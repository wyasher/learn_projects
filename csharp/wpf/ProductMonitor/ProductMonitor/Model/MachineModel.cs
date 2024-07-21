using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ProductMonitor.Model
{
    public class MachineModel
    {
        public string MachineName { get; set; }

        public string Status { get; set; }

        public uint PlanCount { get; set; }

        public uint FinishedCount { get; set; }

        public string OrderNo { get; set; }

        public double Percent => FinishedCount *100.0 /PlanCount;


    }
}
