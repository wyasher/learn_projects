using System;
using System.Linq;
using System.Reflection;

namespace FrameworkDesign.Framework.Singleton
{
    public class Singleton<T> where T:Singleton<T>
    {
        private static T _instance;

        public static T Instance
        {
            get
            {
                if (_instance != null) return _instance;
                // 获取无惨构造函数
                var type = typeof(T);
                var constructors = type.GetConstructors(BindingFlags.Instance | BindingFlags.NonPublic);
                var constructor = Array.Find(constructors, c => c.GetParameters().Length == 0);
                if (constructor == null)
                {
                    throw new Exception("没有无参构造函数" + type.Name);
                }

                _instance = constructor.Invoke(null) as T;

                return _instance;
            }

        }
    }
}