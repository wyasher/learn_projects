using FrameworkDesign.Framework.Architecture.Rule;

namespace FrameworkDesign.Framework.Architecture
{
    public interface IController : IBelongToArchitecture, ICanSendCommand, ICanGetSystem, ICanGetModel
    {
    }
}