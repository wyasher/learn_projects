using System;

namespace FrameworkDesign.Framework
{
    public class Event<T> where T:Event<T>
    {
        private static Action _onEvent;

        public static void Register(Action onEvent)
        {
            _onEvent += onEvent;
        }

        public static void UnRegister(Action onEvent)
        {
            _onEvent -= onEvent;
        }

        public static void Trigger()
        {
            _onEvent?.Invoke();
        }
    }
}