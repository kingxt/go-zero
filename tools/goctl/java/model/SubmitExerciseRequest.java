// Code generated by goctl. DO NOT EDIT.
package com.xhb.logic.http.packet.calcexercise.model;

import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;
import com.xhb.core.packet.HttpData;

public class SubmitExerciseRequest extends HttpData {
	private String unitId = "";
	private int rightCnt;
	private int totalCnt;
	private java.util.ArrayList<QuestionRecord> questions;

	@NotNull 
	public String getUnitId() {
		return this.unitId;
	}

	public void setUnitId(@NotNull String unitId) {
		this.unitId = unitId;
	}

	
	public int getRightCnt() {
		return this.rightCnt;
	}

	public void setRightCnt(int rightCnt) {
		this.rightCnt = rightCnt;
	}

	
	public int getTotalCnt() {
		return this.totalCnt;
	}

	public void setTotalCnt(int totalCnt) {
		this.totalCnt = totalCnt;
	}

	@NotNull 
	public java.util.ArrayList<QuestionRecord> getQuestions() {
		return this.questions;
	}

	public void setQuestions(@NotNull java.util.ArrayList<QuestionRecord> questions) {
		this.questions = questions;
	}
}
