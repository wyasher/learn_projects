using FrameworkDesign.Framework;

namespace FrameworkDesign.Example.Scripts.Model
{
    public interface IGameModel
    {

        public BindableProperty<int> KillCount { get; }

        public  BindableProperty<int> Gold { get; }
        public  BindableProperty<int> Score { get; }
        public  BindableProperty<int> BestScore { get; }
    }
}