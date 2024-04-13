using System;
using System.Collections.Generic;
using FrameworkDesign.Framework.Command;
using FrameworkDesign.Framework.IOC;

namespace FrameworkDesign.Framework.Architecture
{
    public interface IArchitecture
    {
        void RegisterSystem<T>(T system) where T : ISystem;
        T GetUtility<T>() where T : class,IUtility;
        void RegisterModel<T>(T model) where T : IModel;
        T GetSystem<T>() where T : class, ISystem;
        T GetModel<T>() where T : class,IModel;

        void RegisterUtility<T>(T utility) where T : IUtility;

        void SendCommand<T>() where T : ICommand, new();
        void SendCommand<T>(T command) where T : ICommand;
    }

    public abstract class Architecture<T> : IArchitecture where T : Architecture<T>, new()
    {
        private static T _architecture;
        private readonly IOCContainer _iocContainer = new();

        private bool _inited = false;

        private readonly List<IModel> _models = new();
        private readonly List<ISystem> _systems = new();

        private static Action<T> OnRegisterPatch = _architecture => { };

        public static IArchitecture Interface
        {
            get
            {
                if (_architecture == null)
                {
                    MakeSureArchitecture();
                }
                return _architecture;
            }
        }

        private static void MakeSureArchitecture()
        {
            if (_architecture != null) return;
            _architecture = new T();
            _architecture.Init();
            OnRegisterPatch?.Invoke(_architecture);
            foreach (var model in _architecture._models)
            {
                model.Init();
            }
            _architecture._models.Clear();
            foreach (var system in _architecture._systems)
            {
                system.Init();
            }
            _architecture._systems.Clear();
            _architecture._inited = true;
        }

        protected abstract void Init();

        public static T Get<T>() where T : class
        {
            MakeSureArchitecture();
            return _architecture._iocContainer.Get<T>();
        }

        public void Register<T>(T instance)
        {
            MakeSureArchitecture();
            _architecture._iocContainer.Register<T>(instance);
        }

        public void RegisterModel<T>(T model) where T : IModel
        {
            model.SetArchitecture(this);
            _iocContainer.Register(model);
            if (_inited)
            {
                model.Init();
            }
            else
            {
                _models.Add(model);
            }
        }

        public T GetSystem<T>() where T : class, ISystem
        {
            return _iocContainer.Get<T>();
        }

        public T GetModel<T>() where T : class, IModel
        {
            return _iocContainer.Get<T>();
        }

        public void RegisterUtility<T1>(T1 utility) where T1 : IUtility
        {
            _iocContainer.Register<T1>(utility);
        }

        public void SendCommand<T1>() where T1 : ICommand, new()
        {
            var command = new T1();
            command.SetArchitecture(this);
            command.Execute();
        }

        public void SendCommand<T1>(T1 command) where T1 : ICommand
        {
            command.SetArchitecture(this);
            command.Execute();
        }

        public void RegisterSystem<T1>(T1 system) where T1 : ISystem
        {
            system.SetArchitecture(this);
            _iocContainer.Register(system);
            if (_inited)
            {
                system.Init();
            }
            else
            {
                _systems.Add(system);
            }
            
        }

        public T1 GetUtility<T1>() where T1 : class, IUtility
        {
            return _iocContainer.Get<T1>();
        }
    }
}