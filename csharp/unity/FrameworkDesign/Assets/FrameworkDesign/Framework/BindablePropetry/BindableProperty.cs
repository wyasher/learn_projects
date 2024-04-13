using System;

namespace FrameworkDesign.Framework
{
    public class BindableProperty<T> where T : IEquatable<T>
    {
        private T _value;

        public T Value
        {
            get => _value;
            set
            {
                if (_value.Equals(value)) return;
                _value = value;
                OnValueChange?.Invoke(_value);
            }
        }

        public Action<T> OnValueChange;
    }
}