using FrameworkDesign.Example.Scripts.Event;
using UnityEngine;

namespace FrameworkDesign.Example.Scripts.UI
{
    public class UI : MonoBehaviour
    {
        private void Start()
        {
            GamePassEvent.Register(OnGamePass);
        }

        private void OnGamePass()
        {
            transform.Find("Canvas/GamePassPanel").gameObject.SetActive(true);
        }

        private void OnDestroy()
        {
            GameStartEvent.UnRegister(OnGamePass);
        }
    }
}