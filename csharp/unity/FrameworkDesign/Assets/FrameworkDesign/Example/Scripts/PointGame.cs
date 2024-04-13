using FrameworkDesign.Example.Scripts.Model;
using FrameworkDesign.Framework.Architecture;

namespace FrameworkDesign.Example.Scripts
{
    public class PointGame:Architecture<PointGame>
    {
        protected override void Init()
        {
            Register<IGameModel>(new GameModel());
        }
    }
}