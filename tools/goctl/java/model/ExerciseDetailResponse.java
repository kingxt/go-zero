// Code generated by goctl. DO NOT EDIT.
package com.xhb.logic.http.packet.calcexercise.model;

import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;
import com.xhb.core.packet.HttpData;
import com.xhb.core.response.HttpResponseData;

public class ExerciseDetailResponse extends HttpResponseData {
	private ExerciseBrief exerciseBrief;
	private java.util.ArrayList<QuestionDetail> questions;

	@NotNull 
	public ExerciseBrief getExerciseBrief() {
		return this.exerciseBrief;
	}

	public void setExerciseBrief(@NotNull ExerciseBrief exerciseBrief) {
		this.exerciseBrief = exerciseBrief;
	}

	@NotNull 
	public java.util.ArrayList<QuestionDetail> getQuestions() {
		return this.questions;
	}

	public void setQuestions(@NotNull java.util.ArrayList<QuestionDetail> questions) {
		this.questions = questions;
	}
}
