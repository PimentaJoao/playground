#include <opencv2/core.hpp>
#include <opencv2/imgcodecs.hpp>
#include <opencv2/highgui.hpp>
#include <opencv2/features2d.hpp>

#include <iostream>
 
#define CAM_ID 0
#define RATIO_THRESHOLD 0.7f

using namespace cv;
 
int main()
{
    Mat frame_1, frame_2, descriptors_1, descriptors_2, out_frame;

    VideoCapture cap;

    cap.open(CAM_ID, CAP_ANY);

    if (!cap.isOpened()) {
        std::cout << "Unable to open camera!" << std::endl;
        return -1;
    }

    std::vector<KeyPoint> keypoints_1, keypoints_2;

    Ptr<FeatureDetector> detector = ORB::create();
    Ptr<DescriptorExtractor> descriptor = ORB::create();
    Ptr<DescriptorMatcher> matcher = DescriptorMatcher::create(DescriptorMatcher::BRUTEFORCE_HAMMING);

    cap.read(frame_1);
    if (frame_1.empty()) {
        std::cout << "Blank frame 1!" << std::endl;
        return -1;
    }

    while(1) {
        cap.read(frame_2);
        if (frame_2.empty()) {
            std::cout << "Blank frame 2!" << std::endl;
            return -1;
        }

        detector->detectAndCompute(frame_1, noArray(), keypoints_1, descriptors_1);
        detector->detectAndCompute(frame_2, noArray(), keypoints_2, descriptors_2);

        std::vector<std::vector<cv::DMatch>> knn_matches;
        std::vector<DMatch> good_matches;

        matcher->knnMatch(descriptors_1, descriptors_2, knn_matches, 2);

        for (int i = 0; i < knn_matches.size(); i++) {
            if (knn_matches[i][0].distance < RATIO_THRESHOLD * knn_matches[i][1].distance) {
                good_matches.push_back(knn_matches[i][0]);
            }
        }

        drawMatches(frame_1, keypoints_1, frame_2, keypoints_2, good_matches, out_frame, Scalar::all(-1), Scalar::all(-1), std::vector<char>(), DrawMatchesFlags::DEFAULT);

        imshow("Live", out_frame);
        if (waitKey(5) >= 0) {
            break;
        }

        frame_1 = frame_2;
    }

    return 0;
}
