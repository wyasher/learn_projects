namespace FrameworkDesign.Framework.Architecture.Rule
{
    public interface ICanGetModel:IBelongToArchitecture
    {
        
    }
    public static class CanGetModelExtension
    { 
        public static T GetMode<T>(this ICanGetModel self) where T : class, IModel
        {
            return self.GetArchitecture().GetModel<T>();
        }
    }
}