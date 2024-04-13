using FrameworkDesign.Example.Scripts.Command;
using FrameworkDesign.Framework.Architecture;
using UnityEngine;
using UnityEngine.UI;

namespace FrameworkDesign.Example.Scripts.UI
{
    
    public class GameStartPanel : MonoBehaviour,IController
    {
        private void Start()
        {
            transform.Find("BtnStart").GetComponent<Button>().onClick
                .AddListener(() =>
                { 
                    gameObject.SetActive(false);
                    
                    GetArchitecture().SendCommand<StartGameCommand>();
                });
        }

        public IArchitecture GetArchitecture()
        {
            return PointGame.Interface;
        }
    }
}