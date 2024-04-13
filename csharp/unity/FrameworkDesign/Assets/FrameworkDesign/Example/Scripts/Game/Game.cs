using FrameworkDesign.Example.Scripts.Event;
using UnityEngine;

namespace FrameworkDesign.Example.Scripts.Game
{
    public class Game : MonoBehaviour
    {
        private void Awake()
        {
            GameStartEvent.Register(OnGameStart);
        }



        private void OnGameStart() 
        {
            transform.Find("Emenies").gameObject.SetActive(true);
        }

        private void OnDestroy()
        {
            GameStartEvent.UnRegister(OnGameStart);
        }
    }
}