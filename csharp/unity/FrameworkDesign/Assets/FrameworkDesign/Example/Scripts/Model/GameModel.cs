using FrameworkDesign.Framework;
using FrameworkDesign.Framework.Architecture;

namespace FrameworkDesign.Example.Scripts.Model
{
    public class GameModel:AbstractModel,IGameModel
    {
        public BindableProperty<int> KillCount { get; } = new()
        {
            Value = 0
        };
        public BindableProperty<int> Gold { get; }= new()
        {
            Value = 0
        };
        public BindableProperty<int> Score { get; }= new()
        {
            Value = 0
        };
        public BindableProperty<int> BestScore { get; }= new()
        {
            Value = 0
        };

        protected override void OnInit()
        {
            
        }
    }
}