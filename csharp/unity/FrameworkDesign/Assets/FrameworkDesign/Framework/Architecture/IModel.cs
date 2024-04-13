using FrameworkDesign.Framework.Architecture.Rule;

namespace FrameworkDesign.Framework.Architecture
{
    public interface IModel : IBelongToArchitecture, ICanSetArchitecture,ICanGetUtility
    {
        void Init();
    }

    public abstract class AbstractModel : IModel
    {
        private IArchitecture _architecture;

        IArchitecture IBelongToArchitecture.GetArchitecture()
        {
            return _architecture;
        }

        void ICanSetArchitecture.SetArchitecture(IArchitecture architecture)
        {
            _architecture = architecture;
        }

        void IModel.Init()
        {
            OnInit();
        }

        protected abstract void OnInit();
    }
}