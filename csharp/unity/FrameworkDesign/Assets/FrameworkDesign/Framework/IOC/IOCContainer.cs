using System;
using System.Collections.Generic;

namespace FrameworkDesign.Framework.IOC
{
    public class IOCContainer
    {
        private Dictionary<Type, object> _instances = new();

        public void Register<T>(T instance)
        {
            var key = typeof(T);

            if (!_instances.ContainsKey(key))
            {
                _instances.Add(key,instance);
            }
            else
            {
                _instances[key] = instance;
            }
        }

        public T Get<T>() where T : class
        {
            var key = typeof(T);
            if (_instances.TryGetValue(key,out var instance))
            {
                return instance as T;
            }

            return null;
        }

    }
}