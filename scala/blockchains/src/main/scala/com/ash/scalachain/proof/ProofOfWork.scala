package com.ash.scalachain.proof

import com.ash.scalachain.crypto.Crypto

import scala.annotation.tailrec
import spray.json.*
import spray.json.DefaultJsonProtocol.*

object ProofOfWork {
  def proofOfWork(lastHash: String): Long = {
    //    尾随递归优化
    @tailrec
    def prowHelper(lastHash: String, proof: Long): Long = {
      if (validProof(lastHash, proof))
        proof
      else
        prowHelper(lastHash, proof + 1)
    }

    val proof = 0
    prowHelper(lastHash, proof)
  }

  def validProof(lastHash: String, proof: Long): Boolean = {
    val guess = (lastHash ++ proof.toString).toJson.toString
    val guessHash = Crypto.sha256Hash(guess)
    (guessHash take 4) == "0000"
  }
}
